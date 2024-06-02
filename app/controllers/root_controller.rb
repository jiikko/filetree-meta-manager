class RootController < ApplicationController
  skip_before_action :require_login, only: %i[index]

  def index
    render :index
  end
end
