<script lang="ts">
    export let icon = ""
    export let menu:HTMLElement

    const openMenu = () => {
        menu.classList.add("right-0")
        menu.classList.remove("-right-56")
    }
</script>

<div class="text-center">

        {#if icon === "chat" }
            <a href="/">
                <span class="block pb-1">
                    <i class="fa-solid fa-message fa-2xl"></i>
                </span>
                Chat
            </a>
        {:else if icon === "votes" }
                <a
                   class="voteNav"
                   href=""
                   on:click={openMenu}
                   role="button">
                    <span class="block pb-1 voteNav">
                        <i class="fa-solid fa-star fa-2xl voteNav"></i>
                    </span>
                    Votes
                </a>
        {:else if icon === "results" }
            <a href="/results">
                <span class="block pb-1">
                    <i class="fa-solid fa-medal fa-2xl"></i>
                </span>
                Results
            </a>
        {:else if icon === "settings" }
            <a href="/settings">
                <span class="block pb-1">
                    <i class="fa-solid fa-gear fa-2xl"></i>
                </span>
                Settings
            </a>
        {/if}
</div>