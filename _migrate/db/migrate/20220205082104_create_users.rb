class CreateUsers < ActiveRecord::Migration[6.1]
  def change
    create_table :users do |t|
      t.string :name,     null: false, index: { unique: true }
      t.string :password, null: false
      t.timestamps
    end
  end
end
