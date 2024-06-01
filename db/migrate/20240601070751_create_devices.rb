class CreateDevices < ActiveRecord::Migration[7.1]
  def change
    create_table :devices do |t|
      t.bigint :user_id, null: false, index: true
      t.string :name, null: false

      t.timestamps
    end
  end
end
