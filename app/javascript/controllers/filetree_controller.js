import { Controller } from "@hotwired/stimulus";

// Connects to data-controller="filetree"
export default class extends Controller {
  static targets = ["directory"];

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
