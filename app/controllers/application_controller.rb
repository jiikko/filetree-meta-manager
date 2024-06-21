class ApplicationController < ActionController::Base
  before_action :require_login

  # NOTE: productionでは環境変数ENABLE_SIGNUPが必要
  def signup_enabled?
    (Rails.env.production? && ENV['SIGNUP_ENABLED'].present?) || Rails.env.local?
  end

  helper_method :signup_enabled?
end
