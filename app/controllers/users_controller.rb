class UsersController < ApplicationController
  skip_before_action :require_login, only: %i[new create]

  def mypage
    @user = current_user
  end

  def new
    @user = User.new
  end

  def create
    unless signup_enabled?
      redirect_to new_user_path, alert: 'Signup is disabled.'
      return
    end

    @user = User.new(user_params)
    if @user.invalid?
      render :new, status: :unprocessable_entity
      return
    end

    ApplicationRecord.transaction do
      @user.save! && @user.api_keys.create!
    end

    login(params[:email], params[:password])
    redirect_to mypage_path, notice: 'User was successfully created.'
  end

  private

  def user_params
    params.require(:user).permit(:email, :password, :password_confirmation)
  end
end
