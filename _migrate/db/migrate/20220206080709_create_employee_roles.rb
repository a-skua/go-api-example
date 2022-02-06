class CreateEmployeeRoles < ActiveRecord::Migration[6.1]
  def change
    create_table :employee_roles do |t|
      t.belongs_to :company_employee, foreign_key: true
      t.belongs_to :company_role,     foreign_key: true
      t.timestamps
    end
  end
end
