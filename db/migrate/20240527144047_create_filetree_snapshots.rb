class CreateFiletreeSnapshots < ActiveRecord::Migration[7.1]
  def change
    create_table :filetree_snapshots do |t|
      t.json :data, null: false
      t.bigint :user_id, null: false, index: true
      t.bigint :device_id, null: false, index: false
      t.integer :revision, null: false, default: 0
      t.index %i[device_id revision], unique: true

      t.timestamps
    end
  end
end
