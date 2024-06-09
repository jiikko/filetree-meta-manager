module ApplicationHelper
  def render_filetree(node)
    content_tag(:ul) do
      node['children'].map do |child|
        has_children = child['children'].present?
        class_name = has_children ? 'directory' : ''
        data_controller_name = child['children'].present? ? 'filetree' : ''
        row_data_attr = { action: 'click->filetree#toggle' }

        content_tag(:li, class: class_name, data: row_data_attr) do
          toggle_mark_if_directory = if has_children
                                       content_tag(:span, '[+]', data: { filetree_target: 'directory' })
                                     else
                                       content_tag(:span, '')
                                     end
          concat(toggle_mark_if_directory)
          concat(child['path'])
          concat(render_filetree(child)) if has_children
        end
      end.join.html_safe
    end
  end
end
