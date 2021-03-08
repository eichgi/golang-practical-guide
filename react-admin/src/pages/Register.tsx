import React, {Component, SyntheticEvent} from 'react';
import './../Login.css';
import axios from 'axios';
import {Redirect} from 'react-router-dom';

class Register extends Component {

    first_name = '';
    last_name = '';
    email = '';
    password = '';
    password_confirm = '';
    state = {
        registered: false,
    }

    submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        const data = {
            first_name: this.first_name,
            last_name: this.last_name,
            email: this.email,
            password: this.password,
            password_confirm: this.password_confirm,
        };

        const response = await axios.post("http://localhost:8000/api/register", data);
        console.log(response);

        this.setState({
            registered: true
        });
    };

    render() {
        if (this.state.registered) {
            return <Redirect to={'/login'}/>;
        }

        return (
            <main className="form-signin">
                <form onSubmit={this.submit}>
                    <h1 className="h3 mb-3 fw-normal">Please register</h1>

                    <input type="text"
                           className="form-control"
                           placeholder="First Name"
                           required
                           onChange={e => this.first_name = e.target.value}/>

                    <input type="text"
                           className="form-control"
                           placeholder="Last Name"
                           required
                           onChange={e => this.last_name = e.target.value}/>

                    <input type="email"
                           id="inputEmail"
                           className="form-control"
                           placeholder="Email address"
                           required
                           onChange={e => this.email = e.target.value}/>

                    <input type="password"
                           className="form-control"
                           placeholder="Password"
                           required
                           onChange={e => this.password = e.target.value}/>

                    <input type="password"
                           className="form-control"
                           placeholder="Password confirm"
                           required
                           onChange={e => this.password_confirm = e.target.value}/>

                    <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
                    <p className="mt-5 mb-3 text-muted">&copy; 2017–2021</p>
                </form>
            </main>
        );
    }
}

export default Register;