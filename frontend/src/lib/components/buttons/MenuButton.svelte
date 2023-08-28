<script lang="ts">
    import ResultIcon from "$lib/components/icons/ResultIcon.svelte";
    import SettingIcon from "$lib/components/icons/SettingIcon.svelte";
    import VoteIcon from "$lib/components/icons/VoteIcon.svelte";
    import ChatIcon from "$lib/components/icons/ChatIcon.svelte";

    export let icon = ""
    export let menu:HTMLElement

    const openMenu = () => {
        menu.classList.add("right-0")
        menu.classList.remove("-right-56")
    }
</script>

<div class="text-center">

        {#if icon === "chat" }
            <a class="text-neutral-500 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 [&.active]:text-black/90 dark:[&.active]:text-neutral-400" href="/">
                <span class="[&>svg]:w-10">
                    <ChatIcon />
                </span>
                Chat
            </a>
        {:else if icon === "votes" }
            <div id="voteNav">
                <a
                   class="text-neutral-500 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 [&.active]:text-black/90 dark:[&.active]:text-neutral-400"
                   href="#"
                   on:click={openMenu}
                   role="button">
                    <span class="[&>svg]:w-10">
                        <VoteIcon />
                    </span>
                    Votes
                </a>
            </div>
        {:else if icon === "results" }
            <a class="text-neutral-500 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 [&.active]:text-black/90 dark:[&.active]:text-neutral-400" href="/results">
                <span class="[&>svg]:w-10">
                    <ResultIcon />
                </span>
                Results
            </a>
        {:else if icon === "settings" }
            <a class="text-neutral-500 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 [&.active]:text-black/90 dark:[&.active]:text-neutral-400" href="/settings">
                <span class="[&>svg]:w-10">
                    <SettingIcon />
                </span>
                Settings
            </a>
        {/if}
</div>