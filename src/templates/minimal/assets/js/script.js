// Minimal theme toggle
(function() {
  const toggle = document.getElementById('theme-toggle');
  const icon = toggle?.querySelector('.theme-icon');
  const body = document.body;
  
  // Get saved theme or default to dark
  const savedTheme = localStorage.getItem('theme');
  const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
  const currentTheme = savedTheme || (prefersDark ? 'dark' : 'light');
  
  // Apply theme
  if (currentTheme === 'dark') {
    body.classList.add('dark');
    if (icon) icon.textContent = '●';
  } else {
    body.classList.remove('dark');
    if (icon) icon.textContent = '○';
  }
  
  // Toggle handler
  toggle?.addEventListener('click', () => {
    const isDark = body.classList.toggle('dark');
    localStorage.setItem('theme', isDark ? 'dark' : 'light');
    if (icon) icon.textContent = isDark ? '●' : '○';
  });
  
  // Update year
  const yearEl = document.querySelector('.year');
  if (yearEl) yearEl.textContent = new Date().getFullYear();
})();
