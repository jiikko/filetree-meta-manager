class FiletreeSnapshotsController < ApplicationController
  def show
    @filetree_snapshot = current_user.filetree_snapshots.find(params[:id])
    @device = @filetree_snapshot.device
  end

  def destroy
    filetree_snapshot = current_user.filetree_snapshots.find(params[:id])
    filetree_snapshot.destroy
    redirect_to mypage_path
  end
end
