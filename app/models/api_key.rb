class ApiKey < ApplicationRecord
  belongs_to :user

  before_create :generate_token

  def generate_token
    self.value ||= ::UUID7.generate
  end
end
