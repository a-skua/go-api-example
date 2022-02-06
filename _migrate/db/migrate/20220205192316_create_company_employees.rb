class CreateCompanyEmployees < ActiveRecord::Migration[6.1]
  def change
    create_table :company_employees do |t|
      t.belongs_to :company, foreign_key: true
      t.belongs_to :user,    foreign_key: true
      t.timestamps
    end
  end
end
