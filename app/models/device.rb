class Device < ApplicationRecord
  belongs_to :user
  has_many :filetree_snapshots, dependent: :destroy
end
