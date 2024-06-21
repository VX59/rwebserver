import axios from 'axios';

function LoadForms() {
    const EntryForm = document.createElement("form");
    EntryForm.name = "EntryForm";
    const PswdInput = document.createElement("input");
    PswdInput.type = "password";
    PswdInput.placeholder = "authentication key";
    const submit = document.createElement("submit");
    submit.name = "submit"

    body = document.body;

    EntryForm.appendChild(PswdInput);
    EntryForm.appendChild(submit);

    body.appendChild(EntryForm);
}