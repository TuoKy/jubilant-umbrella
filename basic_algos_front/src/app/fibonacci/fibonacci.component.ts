import { Component, Input, OnInit } from '@angular/core';
//import { FormControl } from '@angular/forms';
import { FormBuilder } from '@angular/forms';


import { FibonacciService } from './fibonacci.service';

@Component({
  selector: 'app-fibonacci',
  templateUrl: './fibonacci.component.html',
  styleUrls: ['./fibonacci.component.css']
})
export class FibonacciComponent implements OnInit {

  result: string ="";

  computeForm = this.formBuilder.group({
    n: 2,
    type: 1
  });

  constructor(
    private utils: FibonacciService,
    private formBuilder: FormBuilder,
  ) { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    console.warn('Your order has been submitted', this.computeForm.value);
    //this.computeForm.reset();
    console.warn(this.computeForm.get("n")?.value)
    this.utils.post(this.computeForm.get("n")?.value).subscribe(resp =>{
      this.result = resp
    });
  }

}
