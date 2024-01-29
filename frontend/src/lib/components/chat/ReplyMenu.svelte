<script lang="ts">
    export let replyButtonHandler:Function
</script>

<div
     on:mouseup={replyButtonHandler}
     class="absolute -top-4 right-2 bg-chat-bubble-you px-2 rounded border border-secondary cursor-pointer">
    <i class="fa-solid fa-reply text-typography-main fa-xs"></i>
</div>
