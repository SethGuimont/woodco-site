// Hamburger toggle + dropdown behavior
(function () {
  const btn = document.querySelector(".nav-toggle");
  const nav = document.getElementById("primary-nav");

  if (btn && nav) {
    btn.addEventListener("click", () => {
      const ex = btn.getAttribute("aria-expanded") === "true";
      btn.setAttribute("aria-expanded", String(!ex));
      nav.classList.toggle("open", !ex);

      // Close any open dropdowns if nav is collapsed
      if (ex) {
        document
          .querySelectorAll('.dropbtn[aria-expanded="true"]')
          .forEach((b) => b.setAttribute("aria-expanded", "false"));
      }
    });
  }

  document.querySelectorAll(".dropbtn").forEach((dBtn) => {
    dBtn.addEventListener("click", () => {
      const open = dBtn.getAttribute("aria-expanded") === "true";
      dBtn.setAttribute("aria-expanded", String(!open));

      // Close other dropdowns
      if (!open) {
        document.querySelectorAll(".dropbtn").forEach((b) => {
          if (b !== dBtn) b.setAttribute("aria-expanded", "false");
        });
      }
    });

    dBtn.addEventListener("keydown", (e) => {
      if (e.key === "Escape") dBtn.setAttribute("aria-expanded", "false");
    });

    const menu = dBtn.nextElementSibling;
    if (menu) {
      menu.addEventListener("keydown", (e) => {
        if (e.key === "Escape") dBtn.setAttribute("aria-expanded", "false");
      });
    }
  });

  // Close dropdowns when clicking outside
  document.addEventListener("click", (e) => {
    if (!e.target.closest(".dropdown")) {
      document
        .querySelectorAll(".dropbtn")
        .forEach((b) => b.setAttribute("aria-expanded", "false"));
    }
  });
})();
