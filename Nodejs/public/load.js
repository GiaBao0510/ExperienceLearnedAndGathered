function showLoader() {
    document.querySelector(".loader").style.display = "block";
  }
  
  function hideLoader() {
    document.querySelector(".loader").style.display = "none";
  }
  
  window.onload = showLoader;
  
  // Ẩn loader sau 2 giây
  setTimeout(hideLoader, 2000);