Rails.application.routes.draw do
  get 'up' => 'rails/health#show', as: :rails_health_check

  resource :users, only: %i[new create]
  resources :filetree_snapshots, only: %i[show destroy]

  get :mypage, to: 'users#mypage'

  get 'login' => 'sessions#new', :as => :login
  post 'login' => 'sessions#create'
  delete 'logout' => 'sessions#destroy', :as => :logout

  root 'root#index'

  namespace :api do
    namespace :v1 do
      resources :filetrees, only: [:create]
    end
  end
end
