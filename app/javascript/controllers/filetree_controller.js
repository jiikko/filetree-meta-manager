import { Controller } from "@hotwired/stimulus";

// Connects to data-controller="filetree"
export default class extends Controller {
  static targets = ["directory"];

  connect() {
    // this.directoryTargets.forEach((toggle) => {
    //   toggle.addEventListener("click", this.toggle.bind(this));
    // });
  }

  toggle(event) {
    event.stopPropagation();

    this._toggle(event.target);
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

    if (directory.classList.contains("open")) {
      element.textContent = "[-]";
    } else {
      element.textContent = "[+]";
    }
  }
  _open(element) {
    const parentDirectory = element.closest(".directory");
    parentDirectory.classList.add("open");

    if (element.classList.contains("open")) {
      element.textContent = "[+]";
    } else {
      element.textContent = "[-]";
    }
  }

  _close(element) {
    const parentDirectory = element.closest(".directory");
    parentDirectory.classList.remove("open");

    if (element.classList.contains("open")) {
      element.textContent = "[-]";
    } else {
      element.textContent = "[+]";
    }
  }
}
