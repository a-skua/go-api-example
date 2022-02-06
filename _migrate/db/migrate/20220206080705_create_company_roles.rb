class CreateCompanyRoles < ActiveRecord::Migration[6.1]
  def change
    create_table :company_roles do |t|
      t.belongs_to :company, foreign_key: true
      t.belongs_to :role,    foreign_key: true
      t.timestamps
    end
  end
end
