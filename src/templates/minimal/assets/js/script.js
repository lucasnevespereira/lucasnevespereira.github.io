// Minimal theme toggle
(function () {
  const toggle = document.getElementById("theme-toggle");
  const icon = toggle?.querySelector(".theme-icon");
  const body = document.body;

  // Get saved theme or default to dark
  const savedTheme = localStorage.getItem("theme");
  const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  const currentTheme = savedTheme || (prefersDark ? "dark" : "light");

  // Apply theme
  if (currentTheme === "dark") {
    body.classList.add("dark");
    if (icon) icon.textContent = "●";
  } else {
    body.classList.remove("dark");
    if (icon) icon.textContent = "○";
  }

  // Toggle handler - supports both click and touch
  if (toggle) {
    const handleToggle = (e) => {
      e.preventDefault();
      e.stopPropagation();
      const isDark = body.classList.toggle("dark");
      try {
        localStorage.setItem("theme", isDark ? "dark" : "light");
      } catch (err) {
        // localStorage might not work in private mode
      }
      if (icon) icon.textContent = isDark ? "●" : "○";
    };

    toggle.addEventListener("click", handleToggle);
    toggle.addEventListener("touchend", handleToggle, { passive: false });
  }

  // Update year
  const yearEl = document.querySelector(".year");
  if (yearEl) yearEl.textContent = new Date().getFullYear();
})();
