from app import app
from flask import Flask, jsonify, send_from_directory
from flask_cors import CORS

cors = CORS(app, resources={r"/api/*": {"origins": "*"}})


@app.route('/')
@app.route('/index')
def index():
    return "Hello, World!"


@app.route('/api/v1/register', methods=['GET', 'POST'])
def register():
    return jsonify({
        'status': 1,
        'id': 1
    })


@app.route('/api/v1/login', methods=['GET', 'POST'])
def login():
    return jsonify({
        'status': 1,
        'id': 1
    })


@app.route('/api/v1/logout', methods=['GET', 'POST'])
def logout():
    return jsonify({
        'status': 1
    })


@app.route('/api/v1/companies', methods=['GET', 'POST'])
def companies():
    return jsonify({
        'status': 1,
        'companies': [
            {
                'id': 1,
                'name': 'Ingosstrkh',
                'rating': 4.3
            },
            {
                'id': 2,
                'name': 'Ingosstrkh',
                'rating': 4.3
            },
            {
                'id': 3,
                'name': 'Ingosstrkh',
                'rating': 4.3
            },
            {
                'id': 4,
                'name': 'Ingosstrkh',
                'rating': 4.3
            },
            {
                'id': 5,
                'name': 'Ingosstrkh',
                'rating': 4.3
            },
        ],
    })
