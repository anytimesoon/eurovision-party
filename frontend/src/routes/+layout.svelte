<script lang="ts">
    import "../app.css";

    let theme:string

    if (typeof window !== "undefined"){
        theme = window.localStorage.getItem("theme")

        if (theme === null) {
            window.localStorage.setItem("theme", "classic")
            theme = window.localStorage.getItem("theme")
        }

        document.querySelector("html")?.setAttribute("data-theme", theme)
    }
</script>


<slot />