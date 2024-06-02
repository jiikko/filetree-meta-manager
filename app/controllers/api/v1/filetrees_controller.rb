class Api::V1::FiletreesController < Api::BaseController
  class UnsetDeviceError < StandardError; end

  def create
    raise(UnsetDeviceError, 'need device name') if (device_name = params[:device]).blank?

    ApplicationRecord.transaction do
      device = current_user.devices.find_or_create_by!(name: device_name)
      device.filetree_snapshots.create!(data: filetree_param)
    end
    render json: { status: 'ok' }
  rescue UnsetDeviceError
    render json: { status: 'ng', message: 'device name is not set' }, status: :bad_request
  end

  private

  def filetree_param
    params.require(:filetree).permit!
  end
end
