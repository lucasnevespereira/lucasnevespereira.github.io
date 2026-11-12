(function () {
  const root = document.documentElement;
  const toggle = document.getElementById("theme-toggle");
  const icon = toggle?.querySelector(".theme-icon");

  let saved;
  try {
    saved = localStorage.getItem("theme");
  } catch (err) {}

  const systemDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  let theme = saved || (systemDark ? "dark" : "light");

  const apply = () => {
    root.setAttribute("data-theme", theme);
    if (icon) icon.textContent = theme === "dark" ? "●" : "○";
  };
  apply();

  toggle?.addEventListener("click", () => {
    theme = theme === "dark" ? "light" : "dark";
    try {
      localStorage.setItem("theme", theme);
    } catch (err) {}
    apply();
  });

  const year = document.querySelector(".year");
  if (year) year.textContent = new Date().getFullYear();
})();
