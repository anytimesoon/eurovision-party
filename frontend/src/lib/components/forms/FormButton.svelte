<script lang="ts">
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import Spinner from "$lib/components/Spinner.svelte";

    export let state:number
    let isDisabled:boolean = true

    $: if (state){
        switch (state) {
            case formButtonState.ENABLED:
                isDisabled = false
                break
            default:
                isDisabled = true
        }
    }

    $: disabledClass = isDisabled ? "bg-gray-500 cursor-not-allowed" : ""
</script>


<button type="submit" disabled={isDisabled} class="{disabledClass}" >
    {#if state === formButtonState.SENDING}
        <Spinner size={"s"} thickness={"s"} />
    {:else}
        <slot/>
    {/if}
</button>