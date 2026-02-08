(function () {
  const toggle = document.getElementById("theme-toggle");
  const icon = toggle?.querySelector(".theme-icon");
  const body = document.body;

  let savedTheme;
  try {
    savedTheme = localStorage.getItem("theme");
  } catch (err) {
    console.log("localStorage not available:", err);
  }

  const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  const currentTheme = savedTheme || (prefersDark ? "dark" : "light");

  if (currentTheme === "dark") {
    body.classList.add("dark");
    document.documentElement.classList.remove("light");
    if (icon) icon.textContent = "●";
  } else {
    body.classList.remove("dark");
    document.documentElement.classList.add("light");
    if (icon) icon.textContent = "○";
  }

  if (toggle) {
    const handleToggle = (e) => {
      e.preventDefault();
      e.stopPropagation();
      const isDark = body.classList.toggle("dark");
      if (isDark) {
        document.documentElement.classList.remove("light");
      } else {
        document.documentElement.classList.add("light");
      }
      try {
        localStorage.setItem("theme", isDark ? "dark" : "light");
      } catch (err) {
        console.log("Could not save to localStorage:", err);
      }
      if (icon) icon.textContent = isDark ? "●" : "○";
    };

    toggle.addEventListener("click", handleToggle);
    toggle.addEventListener("touchend", handleToggle, { passive: false });
  }

  const yearEl = document.querySelector(".year");
  if (yearEl) yearEl.textContent = new Date().getFullYear();
})();
