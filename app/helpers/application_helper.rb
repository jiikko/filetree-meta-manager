module ApplicationHelper
  def render_filetree(node)
    content_tag(:ul) do
      node['children'].map do |child|
        content_tag(:li, class: child['children'].present? ? 'directory' : '') do
          concat(content_tag(:span, child['children'].present? ? '[+]' : '', class: 'toggle'))
          concat(child['path'])

          concat(render_filetree(child)) if child['children'].present?
        end
      end.join.html_safe
    end
  end
end
