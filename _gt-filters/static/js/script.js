const filterContainer = document.querySelector(".filter");
const openBtn = document.getElementById("open-filter-btn");
const closeBtn = document.getElementById("close-filter-btn");

function openMenu() {
  filterContainer.classList.add("active");
}

function closeMenu(e) {
  e.stopPropagation();
  filterContainer.classList.remove("active");
}

openBtn.addEventListener("click", openMenu);
closeBtn.addEventListener("click", closeMenu);

const cdRange = document.getElementById("creation-date");
const cdValue = document.getElementById("creation-date-value");
cdRange.addEventListener("input", (e) => {
  cdValue.textContent = cdRange.value;
});

const faRange = document.getElementById("first-album-date");
const faValue = document.getElementById("first-album-date-value");
faRange.addEventListener("input", (e) => {
  faValue.textContent = faRange.value;
});
