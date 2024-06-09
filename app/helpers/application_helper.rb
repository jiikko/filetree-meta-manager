module ApplicationHelper
  def render_filetree(node)
    content_tag(:ul) do
      node['children'].map do |child|
        has_children = child['children'].present?
        class_name = has_children ? 'directory' : nil
        row_data_attr = { action: 'click->filetree#toggle', filetree_target: 'directory' } if has_children

        content_tag(:li, class: class_name, data: row_data_attr) do
          concat(child['path'])
          concat(render_filetree(child)) if has_children
        end
      end.join.html_safe
    end
  end
end
