class FiletreeSnapshotsController < ApplicationController
  def show
    @filetree_snapshot = current_user.filetree_snapshots.find(params[:id])
    @device = @filetree_snapshot.device
  end
end
