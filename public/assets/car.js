let slideIndex = 0;

function moveSlide(n) {
  const slides = document.querySelectorAll(".testimonial");
  const totalSlides = slides.length;

  slideIndex += n;

  if (slideIndex < 0) slideIndex = totalSlides - 1;
  if (slideIndex >= totalSlides) slideIndex = 0;

  const slideWidth = slides[0].offsetWidth;
  const carouselSlide = document.querySelector(".carousel-slide");

  carouselSlide.style.transform = `translateX(${-slideIndex * slideWidth}px)`;
}

// Optional: Auto-play the carousel (every 5 seconds)
setInterval(() => moveSlide(1), 5000);
