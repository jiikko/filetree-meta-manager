class FiletreeSnapshot < ApplicationRecord
  belongs_to :device

  before_create :fill_data_hash

  private

  def fill_data_hash
    self.data_hash = Digest::MD5.hexdigest(data.to_json)
  end
end
