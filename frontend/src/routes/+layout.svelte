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
<svelte:head>
    <link rel="preload" as="font" href="/font/OpenMoji-color-glyf_colr_0.woff2" />
    <link rel="preload" as="font" href="/font/OpenMoji-color-glyf_colr_1.woff2" />
    <link rel="preload" as="font" href="/font/edo.ttf" />
</svelte:head>

<slot />