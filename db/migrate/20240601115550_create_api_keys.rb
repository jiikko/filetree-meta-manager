class CreateApiKeys < ActiveRecord::Migration[7.1]
  def change
    create_table :api_keys do |t|
      t.bigint :user_id, null: false
      t.string :value, null: false

      t.index :value, unique: true
      t.timestamps
    end
  end
end
