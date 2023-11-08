// Theme
let dark = false;

function loadTheme() {
    if (localStorage.getItem("theme")) {
        document.documentElement.setAttribute(
            "data-theme",
            localStorage.getItem("theme")
        );
    }
    dark = localStorage.getItem("theme") == "dark";

    document.getElementById("dark-icon").style.display = dark
        ? "none"
        : "block";
    document.getElementById("light-icon").style.display = dark
        ? "block"
        : "none";
}

loadTheme();

document.querySelector(".theme-toggle").addEventListener("click", () => {
    dark = !dark;
    localStorage.setItem("theme", dark ? "dark" : "light");

    loadTheme();
});
