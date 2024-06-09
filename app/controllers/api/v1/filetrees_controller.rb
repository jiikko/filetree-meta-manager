class Api::V1::FiletreesController < Api::BaseController
  class UnsetDeviceError < StandardError; end
  class SameSnapshotError < StandardError; end

  def create
    raise(UnsetDeviceError, 'need device name') if (device_name = params[:device]).blank?

    ApplicationRecord.transaction do
      device = current_user.devices.find_or_create_by!(name: device_name)
      new_snapshot = device.filetree_snapshots.build(data: filetree_param, revision: device.next_revision)
      new_snapshot.fill_data_hash
      raise SameSnapshotError if new_snapshot.exists_same_snapshot?

      new_snapshot.save!
    end
    device.cleanup_old_revisions # TODO: 非同期処理に逃したい

    render json: { status: 'ok' }
  rescue UnsetDeviceError
    render json: { status: 'ng', message: 'device name is not set' }, status: :bad_request
  rescue SameSnapshotError
    Rails.logger.warn('same snapshot')
    # NOTE: 処理としては正常なので、ステータスコードは 200 で返す
    render json: { status: 'ng', message: 'same snapshot' }, status: :ok
  end

  private

  def filetree_param
    params.require(:filetree).permit!
  end
end
