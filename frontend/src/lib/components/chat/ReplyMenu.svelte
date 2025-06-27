<script lang="ts">
    import Reply from "svelte-material-icons/Reply.svelte";

    interface Props {
        replyButtonHandler: Function;
    }

    let { replyButtonHandler }: Props = $props();
</script>

<div onmouseup={replyButtonHandler} class="absolute
                                            -top-4
                                            right-2
                                            bg-primary
                                            py-1 px-2
                                            rounded
                                            border
                                            border-secondary-light
                                            cursor-pointer">
    <Reply />
</div>
