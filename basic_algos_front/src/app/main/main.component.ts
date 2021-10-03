import { Component, OnInit } from '@angular/core';
import { FibonacciService } from '../fibonacci/fibonacci.service';


@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  text: string ='';

  constructor(private utils: FibonacciService) { }

  ngOnInit(): void {
    //this.getPayload();
  }

  getPayload(): void{
    this.utils.get()
    .subscribe(temp => this.text = temp);
  }

}
