<script lang="ts">
    import {onMount} from "svelte";

    export let navName:string = ""
    export let dest:string = ""
    export let menu:HTMLElement
    let element:HTMLElement

    onMount(() => {
        if(navName === "Vote"){
            element.addEventListener('click', () => {
                menu.classList.add("right-0")
                menu.classList.remove("-right-[75%]")
            })
        }
    })

    $: icon = ():string => {
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
    }
</script>

<a href={dest} class="block">
    <div class="text-center">
        <span class="text-typography-nav" class:voteNav={navName === "Vote"} bind:this={element}>
            <span class="block pb-1 text-4xl" class:voteNav={navName === "Vote"} bind:this={element}>
                {icon()}
            </span>
            {navName}
        </span>
    </div>
</a>
