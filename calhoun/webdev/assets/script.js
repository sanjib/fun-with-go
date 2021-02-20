(function () {
  jQuery(document).ready(function ($) {
    setNavItemActiveClass();

    function setNavItemActiveClass() {
      const currentURL = document.documentURI;
      $(".nav-item").each(function () {
        const navItemURL = $(this).children("a").prop("href");
        if (navItemURL === currentURL) {
          $(this).addClass("active");
        }
      });
    }
  });
})();
