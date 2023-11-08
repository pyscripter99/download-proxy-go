// Theme
function loadTheme() {
    if (localStorage.getItem("theme")) {
        document.documentElement.setAttribute(
            "data-theme",
            localStorage.getItem("theme")
        );

        dark = localStorage.getItem("theme") == "dark";
    }
}

loadTheme();

let dark = false;

document.querySelector(".theme-toggle").addEventListener("click", () => {
    dark = !dark;
    localStorage.setItem("theme", dark ? "dark" : "light");

    loadTheme();
});
