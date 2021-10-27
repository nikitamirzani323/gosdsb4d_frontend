<script>
    import Modal from "../components/Modal.svelte"
    let myModal = "";

    const handlePopup = () => {
        myModal = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal.show();
    };
</script>
<div class="container">
    <header class="d-flex flex-wrap align-items-center justify-content-center justify-content-md-between py-3 mb-4 border-bottom">
        <a href="/" class="d-flex align-items-center col-md-3 mb-2 mb-md-0 text-dark text-decoration-none">
            <svg class="bi me-2" width="40" height="32" role="img" aria-label="Bootstrap"><use xlink:href="#bootstrap"/></svg>
        </a>

        <ul class="nav col-12 col-md-auto mb-2 justify-content-center mb-md-0">
            <li><a href="#" class="nav-link px-2 link-secondary">Home</a></li>
            <li><a
                href="#"
                on:click={() => {
                    handlePopup();
                }} 
                class="nav-link px-2 link-dark">Claim Prize</a></li>
            <li><a
                href="#"
                on:click={() => {
                    handlePopup();
                }} 
                class="nav-link px-2 link-dark">About</a></li>
        </ul>

        <div class="col-md-3 text-end">
            <button
                on:click={() => {
                    handlePopup();
                }}  
                type="button" class="btn btn-info me-2">Login</button>
            <button
                on:click={() => {
                    handlePopup();
                }} 
                type="button" class="btn btn-info">Sign-up</button>
        </div>
    </header>
</div>
<Modal
	modal_id="modalentry"
	modal_size="modal-dialog-centered"
	modal_title=""
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <center>
            Authentication Required
		    <br>
		    Users must be located in Europe
        </center>
	</slot:template>
</Modal>