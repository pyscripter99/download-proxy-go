// Theme
let dark = false;

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

document.querySelector(".theme-toggle").addEventListener("click", () => {
    dark = !dark;
    localStorage.setItem("theme", dark ? "dark" : "light");

    loadTheme();
});
