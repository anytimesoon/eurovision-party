<script lang="ts">
    import {onMount} from "svelte";

    interface Props {
        navName?: string;
        dest?: string;
        menu: HTMLElement;
    }

    let { navName = "", dest = "", menu }: Props = $props();
    let element:HTMLElement = $state()

    onMount(() => {
        if(navName === "Vote"){
            element.addEventListener('click', () => {
                menu.classList.add("right-0")
                menu.classList.remove("-right-[75%]")
            })
        }
    })

    let icon = $derived(():string => {
        switch (navName){
            case "Chat":
                return "🗨"
            case "Vote":
                return "🌟"
            case "Results":
                return "🏅"
            case "Settings":
                return "⚙"
            default:
                return ""
        }
    })
</script>

<a href={dest} class="block">
    <div class="text-center">
        <span class="text-typography-nav" class:voteNav={navName === "Vote"} bind:this={element}>
            <span class="block pb-1 text-2xl" class:voteNav={navName === "Vote"} bind:this={element}>
                {icon()}
            </span>
            {navName}
        </span>
    </div>
</a>
