class Device < ApplicationRecord
  belongs_to :user

  has_many :filetree_snapshots, dependent: :destroy

  def cleanup_old_revisions
    filetree_snapshots.order(revision: :desc).offset(FiletreeSnapshot::KEEP_MAX_REVISIONS).pluck(:id).find_in_batches(&:destroy_all)
  end

  def next_revision
    current_revision = filetree_snapshots.maximum(:revision) || 1
    current_revision + 1
  end
end
