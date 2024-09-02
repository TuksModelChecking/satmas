// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {proto} from '../models';
import {mra} from '../models';
import {experiment} from '../models';
import {context} from '../models';

export function CloseReader():Promise<void>;

export function GenerateMRA(arg1:Array<proto.Agent>,arg2:Array<string>):Promise<mra.MRA>;

export function InitialiseReader():Promise<void>;

export function ReadAllExperimentMetadata():Promise<Array<experiment.Metadata>>;

export function RunExperiment(arg1:experiment.TSExperiment):Promise<void>;

export function SaveExperiment(arg1:experiment.TSExperiment):Promise<void>;

export function SetAppContext(arg1:context.Context):Promise<void>;
