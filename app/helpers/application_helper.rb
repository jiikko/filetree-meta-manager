module ApplicationHelper
  def render_filetree(node)
    content_tag(:ul) do
      node['children'].map do |child|
        has_children = child['children'].present?
        class_names = has_children ? ['directory'] : []
        class_names << 'file_movie' if child['path'].match?(/(mp4|mkv|avi|mov|flv|wmv|mpg|mpeg)$/)
        if has_children
          row_data_attr = {
            controller: 'clipboard',
            action: 'click->filetree#toggle',
            filetree_target: 'directory'
          }
        end

        content_tag(:li, class: class_names.join(' '), data: row_data_attr) do
          concat(child['path'])
          concat(render_filetree(child)) if has_children
        end
      end.join.html_safe
    end
  end
end
