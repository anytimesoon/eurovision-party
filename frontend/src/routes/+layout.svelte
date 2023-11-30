<script>
    import "../app.css";

    let theme = localStorage.getItem("theme") || "classic"

    document.querySelector("html")?.setAttribute("data-theme", theme);
</script>


<slot />