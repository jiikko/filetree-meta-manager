module ApplicationHelper
  # ヘルパーメソッドの追加コメント
  # node: ファイルツリーデータのノード
  # level: 現在のノードの深さ（rootは0）
  def render_filetree(node, level = 0)
    content_tag(:ul, class: ['node', "level-#{level}"]) do
      node['children'].map do |child|
        has_children = child['children'].present?
        class_names = has_children ? ['directory'] : []
        class_names << 'file_movie' if child['path'].match?(/(mp4|mkv|avi|mov|flv|wmv|mpg|mpeg)$/)
        class_names << 'file_photo' if child['path'].match?(/(jpg|jpeg|png|gif|bmp|tiff|webp)$/)
        row_data_attr = if has_children
                          { controller: 'clipboard',
                            action: 'click->filetree#toggle',
                            filetree_target: 'directory' }
                        else
                          { filetree_target: 'file' }
                        end

        content_tag(:li, class: class_names.join(' '), data: row_data_attr) do
          concat(child['path'])
          concat(render_filetree(child, level + 1)) if has_children
        end
      end.join.html_safe
    end
  end
end
