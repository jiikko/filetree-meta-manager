class FiletreeSnapshot < ApplicationRecord
  KEEP_MAX_REVISIONS = 3

  belongs_to :device

  before_create :fill_data_hash

  def fill_data_hash
    self.data_hash ||= Digest::MD5.hexdigest(data.to_json)
  end

  def exists_same_snapshot?
    device.filetree_snapshots.exists?(data_hash:)
  end
end
