<script lang="ts">
    export let replyButtonHandler:Function
</script>

<div on:mouseup={replyButtonHandler} class="absolute -top-4 right-2 bg-gray-600 px-2 rounded border border-secondary cursor-pointer">
    <span class="block rotate-90">⤵</span>
</div>
