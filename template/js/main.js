(function ($) {
  "use strict";

  // Dropdown on mouse hover
  $(document).ready(function () {
    function toggleNavbarMethod() {
      if ($(window).width() > 992) {
        $(".navbar .dropdown")
          .on("mouseover", function () {
            $(".dropdown-toggle", this).trigger("click");
          })
          .on("mouseout", function () {
            $(".dropdown-toggle", this).trigger("click").blur();
          });
      } else {
        $(".navbar .dropdown").off("mouseover").off("mouseout");
      }
    }
    toggleNavbarMethod();
    $(window).resize(toggleNavbarMethod);
  });

  // Back to top button
  $(window).scroll(function () {
    if ($(this).scrollTop() > 100) {
      $(".back-to-top").fadeIn("slow");
    } else {
      $(".back-to-top").fadeOut("slow");
    }
  });
  $(".back-to-top").click(function () {
    $("html, body").animate({ scrollTop: 0 }, 1500, "easeInOutExpo");
    return false;
  });

  // Vendor carousel
  $(".vendor-carousel").owlCarousel({
    loop: true,
    margin: 29,
    nav: false,
    autoplay: true,
    smartSpeed: 1000,
    responsive: {
      0: {
        items: 2,
      },
      576: {
        items: 3,
      },
      768: {
        items: 4,
      },
      992: {
        items: 5,
      },
      1200: {
        items: 6,
      },
    },
  });

  // Related carousel
  $(".related-carousel").owlCarousel({
    loop: true,
    margin: 29,
    nav: false,
    autoplay: true,
    smartSpeed: 1000,
    responsive: {
      0: {
        items: 1,
      },
      576: {
        items: 2,
      },
      768: {
        items: 3,
      },
      992: {
        items: 4,
      },
    },
  });

  // Product Quantity
  $(".quantity button").on("click", function () {
    var button = $(this);
    var oldValue = button.parent().parent().find("input").val();
    if (button.hasClass("btn-plus")) {
      var newVal = parseFloat(oldValue) + 1;
    } else {
      if (oldValue > 0) {
        var newVal = parseFloat(oldValue) - 1;
      } else {
        newVal = 0;
      }
    }
    button.parent().parent().find("input").val(newVal);
  });

  //category list
  $(document).ready(function () {
    //$('#category').onload(function(e) {
    //e.preventDefault();
    $.ajax({
      url: "http://localhost:8081/product",
      type: "GET",
      dataType: "json",
      success: function (result) {
        console.log(result);
        result.forEach(function (value) {
          $("#category").append(
            '<a href="" id = "' +
              value.id +
              '" class="nav-item nav-link">' +
              value.name +
              "</a>"
          );
        });
      },
    });

    $.ajax({
      url: "http://localhost:8081/product",
      type: "GET",
      dataType: "json",
      success: function (result) {
        result.forEach(function (category, index) {
          $("#category").append(
            '<a href="" id = "' +
              category.id +
              '" class="nav-item nav-link">' +
              category.name +
              "</a>"
          );
          if (index < 6) {
            let div1 = $("<div/>");
            div1.attr({ class: "col-lg-4 col-md-6 pb-1" });
            let div2 = $("<div/>");
            div2.attr({
              class: "cat-item d-flex flex-column border mb-4",
              style: "padding: 30px;",
            });
            let p = $("<p/>");
            p.attr({
              class: "text-right",
            }).text(category.numberProduct);
            let a = $("<a/>");
            a.attr({
              class: "cat-img position-relative overflow-hidden mb-3",
              href: "",
            });
            let img = $("<img/>");
            img.attr({
              class: "img-fluid",
              src: "img/cat-4.jpg",
            });
            let h5 = $("<h5/>");
            h5.attr({
              class: "font-weight-semi-bold m-0",
            }).text(category.name);
            $(img).appendTo(a);
            $(h5).appendTo(div2);
            $(p).appendTo(div2);
            $(a).appendTo(div2);
            $(div2).appendTo(div1);
            $(div1).appendTo("#cate-show");
            // .appendTo(divTag);
            console.log(div1);
          }
        });
      },
    });
    //});
  });
})(jQuery);
