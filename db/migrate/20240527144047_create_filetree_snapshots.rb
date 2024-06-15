class CreateFiletreeSnapshots < ActiveRecord::Migration[7.1]
  def change
    # NOTE: dataが大きいので、ROW_FORMAT=COMPRESSEDを指定しておく
    create_table :filetree_snapshots, options: 'ROW_FORMAT=COMPRESSED' do |t|
      t.json :data, null: false
      t.string :data_hash, null: false
      t.bigint :device_id, null: false, index: false
      t.integer :revision, null: false, default: 0
      t.index %i[device_id revision], unique: true
      t.index %i[device_id data_hash], unique: true

      t.timestamps
    end
  end
end
