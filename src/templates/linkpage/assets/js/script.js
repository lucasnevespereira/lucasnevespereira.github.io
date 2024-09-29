console.log("scripts loaded");

const yearDate = new Date().getFullYear().toString();
document.querySelector(".year").innerText = yearDate;

document.addEventListener("DOMContentLoaded", () => {
  const themeToggleButton = document.getElementById("theme-toggle");
  const currentTheme =
    localStorage.getItem("theme") ||
    (window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light");

  if (currentTheme) {
    document.body.classList.add(currentTheme);
  }

  themeToggleButton.addEventListener("click", () => {
    document.body.classList.toggle("dark");
    document.body.classList.toggle("light");

    const newTheme = document.body.classList.contains("dark")
      ? "dark"
      : "light";
    localStorage.setItem("theme", newTheme);
  });
});
