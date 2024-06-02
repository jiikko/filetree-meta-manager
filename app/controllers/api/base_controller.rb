class Api::BaseController < ActionController::Base
  before_action :authenticate_request

  skip_before_action :verify_authenticity_token

  attr_reader :current_user

  private

  def authenticate_request
    token = request.headers['Authorization']
    if token.present? && valid_token?(token)
      @current_user = ApiKey.eager_load(:user).find_by(value: token).user
      return if @current_user
    end

    render json: { error: 'Unauthorized' }, status: :unauthorized
  end

  def valid_token?(_)
    true
  end
end
