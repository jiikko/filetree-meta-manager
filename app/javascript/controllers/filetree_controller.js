import { Controller } from "@hotwired/stimulus";

// Connects to data-controller="filetree"
export default class extends Controller {
  static targets = ["directory", "file", "search"];

  connect() {}

  toggle(event) {
    event.stopPropagation();

    if (event.target.classList.contains("directory")) {
      this._toggle(event.target);
    }
  }

  openAll() {
    this.directoryTargets.forEach((directory) => {
      this._open(directory);
    });
  }

  closeAll() {
    this.directoryTargets.forEach((directory) => {
      this._close(directory);
    });
  }

  search(event) {
    const searchText = event.target.value;
    if (searchText == "" || searchText == undefined) {
      this.element
        .querySelectorAll(".hidden")
        .forEach((el) => el.classList.remove("hidden"));
      return;
    }

    requestAnimationFrame(() => {
      const items = document.querySelectorAll(
        "[data-filetree-target=file], [data-filetree-target=directory]"
      );
      items.forEach((item) => {
        item.classList.add("hidden");
      });

      items.forEach((item) => {
        const text = item.textContent.toLowerCase();
        if (text.includes(searchText.toLowerCase())) {
          item.classList.remove("hidden");
          this.showParents(item);
        }
      });
    });
  }

  showParents(item) {
    let parent = item.parentElement;
    while (parent && parent.tagName !== "body") {
      if (parent.tagName === "ul") {
        parent.classList.remove("hidden");
      }
      parent = parent.parentElement;
    }
  }

  _toggle(element) {
    const directory = element.closest(".directory");
    directory.classList.toggle("open");
  }

  _open(element) {
    const parentDirectory = element.closest(".directory");
    parentDirectory.classList.add("open");
  }

  _close(element) {
    const parentDirectory = element.closest(".directory");
    parentDirectory.classList.remove("open");
  }
}
